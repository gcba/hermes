@extends('voyager::master')

@section('page_title', __('voyager.generic.viewing').' '.$dataType->display_name_plural)

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
                        @if (Voyager::can('add_'.$dataType->name))
                            <hr>
                            <form id="messages-detail-compose">
                                <fieldset>
                                    <textarea class="form-control custom-control" name="message" rows="3" minlength="5" required></textarea>
                                </fieldset>
                                <button type="submit" class="btn btn-primary">Enviar</button>
                            </form>
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

            $('#messages-detail-compose').submit(function(e) {
                e.preventDefault();

                const selectedRow = $('#dataTable .row-selected').first();
                const rowData = $('#dataTable').DataTable().row(selectedRow).data();
                const textarea = $('#messages-detail-compose textarea');
                const csrfToken = $('meta[name="csrf-token"]').attr('content');

                fetch('/admin/messages', {
                    method: 'post',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                        'X-CSRF-TOKEN': csrfToken
                    },
                    body: JSON.stringify({
                        message: textarea.val().trim(),
                        rating: rowData.rating_id
                    })
                })
                .then((response) => response.json())
                .then((json) => {
                    if (json.status === 201) {
                        textarea.val('');
                        appendMessage(json.message);
                    }
                    else console.error(json);
                })
                .catch((error) => {
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
                    { data: 'message', name: 'message' },
                    { data: 'rating.app.name', name: 'app' },
                    { data: 'rating.rating', name: 'rating.rating', visible: false },
                    { data: 'created_at', name: 'created_at' }
                ],
                bSort: false,
                bInfo: false,
                mark: true,
                language: {
                    search: '',
                    sLengthMenu: '_MENU_'
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
            .on('stateLoaded.dt', function (e, settings, data) {
                console.log("hola");
            })
            .on('preDraw', function (e, settings) {
                $(this).DataTable().rows().every(function () {
                    if (this.data().status === 0) {
                        $(this.node()).addClass('row-unread');
                    }
                });
            });
        })
        .on('click', 'tbody tr', function() {
            if ($(this).children().length <= 1) return;

            selectRow(this);
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
        }

        const messageHeading = function(content, direction) {
            return $('<div>', {
                class: 'message-pill',
                text: content
             });
        }

        const messageBody = function(content) {
            return $('<div>', {
                class: 'panel-body message-body',
                text: content
             });
        }

        const threadHeading = function(name, row) {
            const container = $('<div>', { class: 'messages-detail-header' });
            const user = $('<h3>', { text: name + ' ' });

            const contextualInfo = $('<small>', {
                text: row.rating.app_version ?
                `${row.rating.app.name} ${row.rating.app_version}, ${row.rating.platform.name}` :
                `${row.rating.app.name}, ${row.rating.platform.name}`
            });

            user.append(contextualInfo);
            container.append(user);

            return container;
        }

        const buildMessage = function(content) {
            const message = messagePanel(content.direction);
            const heading = messageHeading(content.created_at, content.direction);
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
        }

        const buildThread = function(messages, row) {
            const thread = $('.messages-detail-list').first().empty();

            if (row.rating && row.rating.appuser) {
                thread.append(threadHeading(row.rating.appuser.name, row));
            }
            else {
                thread.append(threadHeading('Anónimo', row));
            }

            for (const message of messages) {
               thread.append(buildMessage(message));
            }
        }

        const selectRow = function(row) {
            const rowData = $('#dataTable').DataTable().row(row).data();
            const $row = $(row);

            $('#dataTable .row-selected').removeClass('row-selected');

            if (!$row.hasClass('row-search')) {
                $row.addClass('row-selected');

                if (rowData) {
                    const ratingID = rowData.rating_id;

                    fetch('/admin/ratings/' + ratingID + '/messages', {
                        method: 'GET',
                        credentials: 'include'
                    })
                    .then(function(response) {
                        return response.json();
                    })
                    .then(function(response) {
                        if ($row.hasClass('row-unread')) {
                            $row.removeClass('row-unread');
                        }

                        buildThread(response, rowData);
                    })
                }
            }
        }

        const appendMessage = function(message) {
            const thread = $('.messages-detail-list').first();

            thread.append(buildMessage(message));
        }
    </script>
@stop