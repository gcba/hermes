@extends('voyager::master')

@section('page_title', __('voyager.generic.viewing') . ' ' . $dataType->display_name_plural)

@section('page_header')
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <h1 class="page-title">
        <i class="{{ $dataType->icon }}"></i> {{ $dataType->display_name_plural }}
    </h1>
    @include('voyager::multilingual.language-selector')
@stop

@section('content')
    <div class="page-content browse container-fluid">
        @include('voyager::alerts')
        <div class="row">

            <div class="col-md-6 messages-master">
                <div class="panel panel-bordered">
                    <div class="panel-body table-responsive">
                        <table id="dataTable" class="row table table-hover">
                            <tbody></tbody>
                            <tfoot>
                                @foreach($dataType->browseRows as $row)
                                    <th></th>
                                @endforeach
                            </tfoot>
                        </table>
                        @if (isset($dataType->server_side) && $dataType->server_side)
                            <div class="pull-left">
                                <div role="status" class="show-res" aria-live="polite">{{ trans_choice(
                                    'voyager.generic.showing_entries', $dataTypeContent->total(), [
                                        'from' => $dataTypeContent->firstItem(),
                                        'to' => $dataTypeContent->lastItem(),
                                        'all' => $dataTypeContent->total()
                                    ]) }}</div>
                            </div>
                            <div class="pull-right">
                                {{ $dataTypeContent->links() }}
                            </div>
                        @endif
                    </div>
                </div>
            </div>

            <div class="col-md-6 messages-detail">
                <div class="panel panel-bordered">
                    <div class="panel-body">
                        <div class="messages-detail-list">
                        </div>
                        @if (Voyager::can('add_' . $dataType->name))
                            <div class="messages-detail-compose hidden">
                                <hr>
                                <form id="messages-form">
                                    <fieldset>
                                        <textarea class="form-control custom-control" name="message" rows="3" minlength="5" maxlength="1500" required></textarea>
                                    </fieldset>
                                    <button type="submit" class="btn btn-primary">Enviar</button>
                                </form>
                            </div>
                        @endif
                    </div>
                </div>
            </div>

        </div>
    </div>
@stop

@section('css')
@if(!$dataType->server_side && config('dashboard.data_tables.responsive'))
<link rel="stylesheet" href="{{ voyager_asset('lib/css/responsive.dataTables.min.css') }}">
@endif
@stop

@section('javascript')
    <!-- DataTables -->
    @if(!$dataType->server_side && config('dashboard.data_tables.responsive'))
        <script src="{{ voyager_asset('lib/js/dataTables.responsive.min.js') }}"></script>
    @endif
    @if($isModelTranslatable)
        <script src="{{ voyager_asset('js/multilingual.js') }}"></script>
    @endif
    <script>
        // TODO: Create a proper asset pipeline

        $(document).ready(function () {
            @if ($isModelTranslatable)
                $('.side-body').multilingual();
            @endif

            $('#messages-form').submit(function(e) {
                e.preventDefault();

                const selectedRow = $('#dataTable .row-selected').first();
                const selectedRowIndex = selectedRow.index();
                const rowData = $('#dataTable').DataTable().row(selectedRow).data();
                const csrfToken = $('meta[name="csrf-token"]').attr('content');
                const errorTitle = 'Error';
                const errorMessage = 'No se pudo enviar el mensaje.';
                const messageText = $('#messages-form textarea').val().trim();

                if (messageText.length === 0) {
                    toastr.warning('El mensaje está vacío.');

                    return false;
                }

                disableForm();

                fetch('/admin/messages', {
                    method: 'post',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                        'X-CSRF-TOKEN': csrfToken
                    },
                    body: JSON.stringify({
                        message: messageText,
                        rating: rowData.id
                    })
                })
                .then((response) => response.json())
                .then((json) => {
                    if (json.status === 201) {
                        reloadThreads(selectedRowIndex);
                        clearForm();
                        appendMessage(json.message);
                    }
                    else {
                        toastr.error(errorMessage, errorTitle);
                        console.error(json);
                    }

                    enableForm();
                })
                .catch((error) => {
                    enableForm();
                    toastr.error(errorMessage, errorTitle);
                    console.error(error);
                });

                return false;
            });

            $('#dataTable').DataTable({
                processing: true,
                serverSide: true,
                ajax: {
                    url: '{!! route('messages.api') !!}',
                    data: function (d) {
                        d.columns.forEach(function (column) {
                            if (column.name && column.name.indexOf('.') != -1) {
                                const name = column.name.replace('.', '_');
                                const searchTerm = $('input[name=' + name + ']').val();

                                if (searchTerm && searchTerm.trim().length > 0) d[name] = searchTerm.trim();
                            }
                        });
                    }
                },
                columns: [
                    { data: 'messages.message', name: 'messages.message' },
                    { data: 'app.name', name: 'app.name' },
                    { data: 'rating', name: 'rating', visible: false },
                    { data: 'messages.created_at', name: 'messages.created_at' }
                ],
                dom: '<"top"i>rt<"bottom"flp><"clear">',
                bSort: false,
                bInfo: false,
                mark: true,
                language: {
                    search: '',
                    sLengthMenu: '_MENU_',
                    sEmptyTable: 'Sin mensajes para mostrar',
                    sLoadingRecords: 'Cargando...',
                    sProcessing: 'Procesando...',
                    sZeroRecords: 'No se encontraron registros coincidentes',
                    oPaginate: {
                        sFirst: 'Primero',
                        sLast: 'Último',
                        sNext: 'Siguiente',
                        sPrevious: 'Anterior'
                    }
                },
                initComplete: function () {
                    if ($('.dataTables_empty').length !== 0) return;

                    this.api().columns().every(function () {
                        const column = this;
                        const input = document.createElement('input');

                        $(input).appendTo($(column.footer()).empty())
                            .on('change', function () {
                                const $this = $(this);
                                const val = $.fn.dataTable.util.escapeRegex($this.val().trim());

                                column.search($this.val()).draw();
                            })
                            .closest('tr').addClass('row-search');
                    });

                    selectRow($('#dataTable tbody tr:nth-child(1)'));
                }
            })
            .on('preDraw', function (e, settings) {
                $(this).DataTable().rows().every(function () {
                    if (this.data().messages.status === 0 && this.data().messages.direction !== 'out') {
                        $(this.node()).addClass('row-unread');
                    }
                });
            });
        })
        .on('click', 'tbody tr', function() {
            if ($(this).children().length <= 1) return;

            selectRow(this);
        });

        $('#search-input select').select2({
            minimumResultsForSearch: Infinity
        });

        var deleteFormAction;

        $('td').on('click', '.delete', function (e) {
            var form = $('#delete_form')[0];

            if (!deleteFormAction) { // Save form action initial value
                deleteFormAction = form.action;
            }

            form.action = deleteFormAction.match(/\/[0-9]+$/)
                ? deleteFormAction.replace(/([0-9]+$)/, $(this).data('id'))
                : deleteFormAction + '/' + $(this).data('id');

            $('#delete_modal').modal('show');
        });

        const messagePanel = function(direction) {
            return $('<div>', { class: 'panel panel-default message message-' + direction });
        };

        const messageHeading = function(id, content, direction) {
            return $('<a>', {
                class: 'message-pill',
                text: content,
                title: 'Ver mensaje',
                href: '/admin/messages/' + id
            });
        };

        const messageBody = function(content) {
            return $('<div>', {
                class: 'panel-body message-body',
                text: content
             });
        };

        const threadHeading = function(name, row) {
            const container = $('<div>', { class: 'messages-detail-header' });
            const user = $('<h3>', { text: name });
            const rating = $('<span>', { text: row.rating, class: 'label label-default' });

            const contextualInfo = $('<small>', {
                text: row.app_version ?
                `${row.app.name} ${row.app_version}, ${row.platform.name}` :
                `${row.app.name}, ${row.platform.name}`
            });

            user.prepend(rating);
            user.append(contextualInfo);
            container.append(user);

            return container;
        };

        const buildMessage = function(content) {
            const message = messagePanel(content.direction);
            const heading = messageHeading(content.id, content.created_at, content.direction);
            const body = messageBody(content.message);

            if (content.direction === 'in') {
                message.append(body);
                message.append(heading);
            }
            else {
                message.append(heading);
                message.append(body);
            }

            return message;
        };

        const buildThread = function(messages, row) {
            const thread = $('.messages-detail-list').first().empty();

            if (row.appuser) {
                thread.append(threadHeading(row.appuser.name, row));
            }
            else {
                thread.append(threadHeading('Anónimo', row));
            }

            for (const message of messages) {
               thread.append(buildMessage(message));
            }
        };

        const selectRow = function(row) {
            const rowData = $('#dataTable').DataTable().row(row).data();
            const isAnonimous = rowData.appuser_id === null || rowData.appuser === null;
            const hasNoEmail = rowData.appuser !== null && rowData.appuser.email === null;
            const $row = $(row);

            if (!$row.hasClass('row-search')) {
                clearForm();

                if (isAnonimous || hasNoEmail) disableForm();

                if (!$row.hasClass('row-selected')) {
                    hideMessages();
                    hideForm();
                }

                $('#dataTable .row-selected').removeClass('row-selected');
                $row.addClass('row-selected');

                if (rowData) {
                    const ratingID = rowData.id;

                    fetch('/admin/ratings/' + ratingID + '/messages', {
                        method: 'GET',
                        credentials: 'include'
                    })
                    .then(function(response) {
                        return response.json();
                    })
                    .then(function(response) {
                        if (!isAnonimous && !hasNoEmail) {
                            enableForm();
                            showForm();
                        }

                        $row.removeClass('row-unread');
                        showMessages();
                        buildThread(response, rowData);
                    })
                }
            }
        };

        const appendMessage = function(message) {
            const thread = $('.messages-detail-list').first();
            const newMessage = buildMessage(message);

            newMessage.hide().appendTo(thread).fadeIn('slow');
        };

        const showMessages = function() {
            $('.messages-detail-list').removeClass('hidden');
        };

        const hideMessages = function() {
            $('.messages-detail-list').addClass('hidden');
        };

        const showForm = function() {
            $('.messages-detail-compose').removeClass('hidden');
        };

        const hideForm = function() {
            $('.messages-detail-compose').addClass('hidden');
        };

        const enableForm = function() {
            $('#messages-form :input, #messages-form button').prop('disabled', false);
        };

        const disableForm = function() {
            $('#messages-form :input, #messages-form button').prop('disabled', true);
        };

        const clearForm = function() {
            $('#messages-form textarea').val('');
        };

        const reloadThreads = function(currentThread) {
            $('#dataTable').DataTable().ajax.reload(() => {
                $('#dataTable tbody').children().eq(currentThread).addClass('row-selected');
            }, false);
        };
    </script>
@stop