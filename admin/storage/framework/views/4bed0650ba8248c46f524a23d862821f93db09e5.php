<?php $__env->startSection('css'); ?>
    <script type="text/javascript" src="<?php echo e(voyager_asset('js/vue21.min.js')); ?>"></script>
    <link rel="stylesheet" href="<?php echo e(voyager_asset('css/database.css')); ?>">
<?php $__env->stopSection(); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="voyager-data"></i> Database
        <a href="<?php echo e(route('voyager.database.create')); ?>" class="btn btn-success"><i class="voyager-plus"></i>
            Create New Table</a>
    </h1>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>

    <div class="page-content container-fluid">
        <?php echo $__env->make('voyager::alerts', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <div class="row">
            <div class="col-md-12">

                <table class="table table-striped database-tables">
                    <thead>
                        <tr>
                            <th>Table Name</th>
                            <th>BREAD/CRUD Actions</th>
                            <th style="text-align:right">Table Actions</th>
                        </tr>
                    </thead>

            <?php $__currentLoopData = $tables; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $table): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                    <?php if(in_array($table->name, config('voyager.database.tables.hidden', []))) continue; ?>
                    <tr>
                        <td>
                            <p class="name">
                                <a href="<?php echo e(route('voyager.database.show', $table->name)); ?>"
                                   data-name="<?php echo e($table->name); ?>" class="desctable">
                                   <?php echo e($table->name); ?>

                                </a>
                            <?php if($table->dataTypeId): ?>
                                <i class="voyager-bread"
                                   style="font-size:25px; position:absolute; margin-left:10px; margin-top:-3px;"></i>
                            <?php endif; ?>
                            </p>
                        </td>

                        <td>
                            <div class="bread_actions">
                            <?php if($table->dataTypeId): ?>
                                <a href="<?php echo e(route('voyager.database.bread.edit', $table->name)); ?>"
                                   class="btn-sm btn-default edit">
                                   Edit BREAD
                                </a>
                                <div data-id="<?php echo e($table->dataTypeId); ?>" data-name="<?php echo e($table->name); ?>"
                                     class="btn-sm btn-danger delete" style="display:inline">
                                     Delete BREAD
                                </div>
                            <?php else: ?>
                                <a href="<?php echo e(route('voyager.database.bread.create', ['name' => $table->name])); ?>"
                                   class="btn-sm btn-default">
                                    <i class="voyager-plus"></i> Add BREAD to this table
                                </a>
                            <?php endif; ?>
                            </div>
                        </td>

                        <td class="actions">
                            <a class="btn-danger btn-sm pull-right delete_table <?php if($table->dataTypeId): ?> remove-bread-warning <?php endif; ?>"
                               data-table="<?php echo e($table->name); ?>" style="display:inline; cursor:pointer;">
                               <i class="voyager-trash"></i> Delete
                            </a>
                            <a href="<?php echo e(route('voyager.database.edit', $table->name)); ?>"
                               class="btn-sm btn-primary pull-right" style="display:inline; margin-right:10px;">
                               <i class="voyager-edit"></i> Edit
                            </a>
                            <a href="<?php echo e(route('voyager.database.show', $table->name)); ?>"
                               data-name="<?php echo e($table->name); ?>"
                               class="btn-sm btn-warning pull-right desctable" style="display:inline; margin-right:10px;">
                               <i class="voyager-eye"></i> View
                            </a>
                        </td>
                    </tr>
                <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                </table>
            </div>
        </div>
    </div>

    <div class="modal modal-danger fade" tabindex="-1" id="delete_builder_modal" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title"><i class="voyager-trash"></i> Are you sure you want to delete the BREAD for
                        the <span id="delete_builder_name"></span> table?</h4>
                </div>
                <div class="modal-footer">
                    <form action="<?php echo e(route('voyager.database.bread.delete', ['id' => null])); ?>" id="delete_builder_form" method="POST">
                        <?php echo e(method_field('DELETE')); ?>

                        <input type="hidden" name="_token" value="<?php echo e(csrf_token()); ?>">
                        <input type="submit" class="btn btn-danger" value="Yes, remove the BREAD">
                    </form>
                    <button type="button" class="btn btn-outline pull-right" data-dismiss="modal">Cancel</button>
                </div>
            </div><!-- /.modal-content -->
        </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->

    <div class="modal modal-danger fade" tabindex="-1" id="delete_modal" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title"><i class="voyager-trash"></i> Are you sure you want to delete the <span
                                id="delete_table_name"></span> table?</h4>
                </div>
                <div class="modal-footer">
                    <form action="<?php echo e(route('voyager.database.destroy', ['database' => '__database'])); ?>" id="delete_table_form" method="POST">
                        <?php echo e(method_field('DELETE')); ?>

                        <input type="hidden" name="_token" value="<?php echo e(csrf_token()); ?>">
                        <input type="submit" class="btn btn-danger pull-right" value="Yes, delete this table">
                        <button type="button" class="btn btn-outline pull-right" style="margin-right:10px;"
                                data-dismiss="modal">Cancel
                        </button>
                    </form>

                </div>
            </div><!-- /.modal-content -->
        </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->

    <div class="modal modal-info fade" tabindex="-1" id="table_info" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title"><i class="voyager-data"></i> {{ table.name }}</h4>
                </div>
                <div class="modal-body" style="overflow:scroll">
                    <table class="table table-striped">
                        <thead>
                        <tr>
                            <th>Field</th>
                            <th>Type</th>
                            <th>Null</th>
                            <th>Key</th>
                            <th>Default</th>
                            <th>Extra</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr v-for="row in table.rows">
                            <td><strong>{{ row.Field }}</strong></td>
                            <td>{{ row.Type }}</td>
                            <td>{{ row.Null }}</td>
                            <td>{{ row.Key }}</td>
                            <td>{{ row.Default }}</td>
                            <td>{{ row.Extra }}</td>
                        </tr>
                        </tbody>
                    </table>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-outline pull-right" data-dismiss="modal">Close</button>
                </div>
            </div><!-- /.modal-content -->
        </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->

<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>

    <script>

        var table = {
            name: '',
            rows: []
        };

        new Vue({
            el: '#table_info',
            data: {
                table: table,
            },
        });

        $(function () {

            $('.bread_actions').on('click', '.delete', function (e) {
                id = $(this).data('id');
                name = $(this).data('name');

                $('#delete_builder_name').text(name);
                $('#delete_builder_form')[0].action += '/' + id;
                $('#delete_builder_modal').modal('show');
            });

            $('.database-tables').on('click', '.desctable', function (e) {
                e.preventDefault();
                href = $(this).attr('href');
                table.name = $(this).data('name');
                table.rows = [];
                $.get(href, function (data) {
                    $.each(data, function (key, val) {
                        table.rows.push({
                            Field: val.field,
                            Type: val.type,
                            Null: val.null,
                            Key: val.key,
                            Default: val.default,
                            Extra: val.extra
                        });
                        $('#table_info').modal('show');
                    });
                });
            });

            $('td.actions').on('click', '.delete_table', function (e) {
                table = $(this).data('table');
                if ($(this).hasClass('remove-bread-warning')) {
                    toastr.warning("Please make sure to remove the BREAD on this table before deleting the table.");
                } else {
                    $('#delete_table_name').text(table);
                    $('#delete_table_form')[0].action = $('#delete_table_form')[0].action.replace('__database', table);
                    $('#delete_modal').modal('show');
                }
            });

        });
    </script>

<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>