@extends('voyager::master')

@section('page_title', __('voyager.generic.viewing').' '.$dataType->display_name_plural)

@section('page_header')
    <div class="container-fluid">
        <h1 class="page-title">
            <i class="{{ $dataType->icon }}"></i> {{ $dataType->display_name_plural }}
        </h1>
        @include('voyager::multilingual.language-selector')
    </div>
@stop

@section('content')
    <div class="page-content browse container-fluid">
        @include('voyager::alerts')
        <div class="row">
            <div class="col-md-12">
                <div class="panel panel-bordered">
                    <div class="panel-body table-responsive">
                        <table id="dataTable" class="table table-hover table-ratings">
                            <thead>
                                <tr>
                                    @foreach($dataType->browseRows as $row)
                                    <th>
                                        {{ $row->display_name }}
                                    </th>
                                    @endforeach
                                    <th class="actions">Ver</th>
                                </tr>
                            </thead>
                            <tbody>
                            </tbody>
                            <tfoot>
                                @foreach($dataType->browseRows as $row)
                                    <th></th>
                                @endforeach
                            </tfoot>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>

    {{-- Single delete modal --}}
    <div class="modal modal-danger fade" tabindex="-1" id="delete_modal" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="{{ __('voyager.generic.close') }}"><span
                                aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title"><i class="voyager-trash"></i> {{ __('voyager.generic.delete_question') }} {{ strtolower($dataType->display_name_singular) }}?</h4>
                </div>
                <div class="modal-footer">
                    <form action="{{ route('voyager.'.$dataType->slug.'.index') }}" id="delete_form" method="POST">
                        {{ method_field("DELETE") }}
                        {{ csrf_field() }}
                        <input type="submit" class="btn btn-danger pull-right delete-confirm"
                                 value="{{ __('voyager.generic.delete_confirm') }} {{ strtolower($dataType->display_name_singular) }}">
                    </form>
                    <button type="button" class="btn btn-default pull-right" data-dismiss="modal">{{ __('voyager.generic.cancel') }}</button>
                </div>
            </div><!-- /.modal-content -->
        </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->
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
        $(document).ready(function () {
            @if ($isModelTranslatable)
                $('.side-body').multilingual();
            @endif

            $('#dataTable').DataTable({
                processing: true,
                serverSide: true,
                language: {!! json_encode(["language" => __('voyager.datatable')]) !!}.language,
                ajax: {
                    url: '{!! route('ratings.api') !!}',
                    data: function (d) {
                        d.columns.forEach(function (column) {
                            if (column.name && column.name.indexOf('.') != -1) {
                                var name = column.name.replace('.', '_');
                                var searchTerm = $('input[name=' + name + ']').val();

                                if (searchTerm && searchTerm.trim().length > 0) d[name] = searchTerm.trim();
                            }
                        });
                    }
                },
                columns: [
                    { data: 'rating', name: 'rating' },
                    { data: 'range.name', name: 'range.name' },
                    { data: 'description', name: 'description' },
                    { data: 'has_message', name: 'has_message' },
                    { data: 'app.name', name: 'app.name' },
                    { data: 'app_version', name: 'app_version' },
                    { data: 'platform.name', name: 'platform.name' },
                    { data: 'platform_version', name: 'platform_version' },
                    { data: 'browser_id', name: 'browser_id' },
                    { data: 'browser_version', name: 'browser_version' },
                    { data: 'appuser.name', name: 'appuser.name' },
                    { data: 'device.name', name: 'device.name' },
                    { data: 'created_at', name: 'created_at' },
                    {
                        'data': null,
                        'render': function (data, type, row, meta) {
                            return '<a href="/admin/ratings/' + data.id + '" title="Ver" class="btn btn-sm btn-warning pull-right"><i class="voyager-eye"></i><span class="hidden-xs hidden-sm"></span></a>';
                        },
                        'targets': -1
                    }
                ],
                order: [[12, 'desc']],
                bFilter: false,
                mark: true,
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
                }
            });
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
            console.log(form.action);

            $('#delete_modal').modal('show');
        });
    </script>
@stop