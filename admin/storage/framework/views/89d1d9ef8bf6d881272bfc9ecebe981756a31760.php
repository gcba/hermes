<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="voyager-list-add"></i> <?php echo e($dataType->display_name_plural); ?>

        <?php if(Voyager::can('add_'.$dataType->name)): ?>
            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.create')); ?>" class="btn btn-success">
                <i class="voyager-plus"></i> Add New
            </a>
        <?php endif; ?>
    </h1>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>
    <?php echo $__env->make('voyager::menus.partial.notice', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>

    <div class="page-content container-fluid">
        <?php echo $__env->make('voyager::alerts', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <div class="row">
            <div class="col-md-12">
                <div class="panel panel-bordered">
                    <div class="panel-body">
                        <table id="dataTable" class="table table-hover">
                            <thead>
                            <tr>
                                <?php $__currentLoopData = $dataType->browseRows; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $rows): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                <th><?php echo e($rows->display_name); ?></th>
                                <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                <th class="actions">Actions</th>
                            </tr>
                            </thead>
                            <tbody>
                                <?php $__currentLoopData = $dataTypeContent; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $data): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                <tr>
                                    <?php $__currentLoopData = $dataType->browseRows; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $row): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                    <td>
                                        <?php if($row->type == 'image'): ?>
                                            <img src="<?php if( strpos($data->{$row->field}, 'http://') === false && strpos($data->{$row->field}, 'https://') === false): ?><?php echo e(Voyager::image( $data->{$row->field} )); ?><?php else: ?><?php echo e($data->{$row->field}); ?><?php endif; ?>" style="width:100px">
                                        <?php else: ?>
                                            <?php echo e($data->{$row->field}); ?>

                                        <?php endif; ?>
                                    </td>
                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                    <td class="no-sort no-click">
                                        <?php if(Voyager::can('delete_'.$dataType->name)): ?>
                                            <div class="btn-sm btn-danger pull-right delete" data-id="<?php echo e($data->id); ?>">
                                                <i class="voyager-trash"></i> Delete
                                            </div>
                                        <?php endif; ?>
                                        <?php if(Voyager::can('edit_'.$dataType->name)): ?>
                                            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.edit', $data->id)); ?>" class="btn-sm btn-primary pull-right edit">
                                                <i class="voyager-edit"></i> Edit
                                            </a>
                                        <?php endif; ?>
                                        <?php if(Voyager::can('edit_'.$dataType->name)): ?>
                                            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.builder', $data->id)); ?>" class="btn-sm btn-success pull-right">
                                                <i class="voyager-list"></i> Builder
                                            </a>
                                        <?php endif; ?>
                                    </td>
                                </tr>
                                <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="modal modal-danger fade" tabindex="-1" id="delete_modal" role="dialog">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    <h4 class="modal-title">
                        <i class="voyager-trash"></i> Are you sure you want to delete this <?php echo e($dataType->display_name_singular); ?>?
                    </h4>
                </div>
                <div class="modal-footer">
                    <form action="<?php echo e(route('voyager.'.$dataType->slug.'.index')); ?>" id="delete_form" method="POST">
                        <?php echo e(method_field("DELETE")); ?>

                        <?php echo e(csrf_field()); ?>

                        <input type="submit" class="btn btn-danger pull-right delete-confirm" value="Yes, Delete This <?php echo e($dataType->display_name_singular); ?>">
                    </form>
                    <button type="button" class="btn btn-default pull-right" data-dismiss="modal">Cancel</button>
                </div>
            </div>
        </div>
    </div>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>
    <!-- DataTables -->
    <script>
        $(document).ready(function () {
            $('#dataTable').DataTable({ "order": [] });
        });

        $('td').on('click', '.delete', function (e) {
            id = $(e.target).data('id');

            $('#delete_form')[0].action += '/' + id;

            $('#delete_modal').modal('show');
        });
    </script>
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>