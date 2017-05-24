<?php $__env->startSection('page_title','All '.$dataType->display_name_plural); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="voyager-news"></i> <?php echo e($dataType->display_name_plural); ?>

        <?php if(Voyager::can('add_'.$dataType->name)): ?>
            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.create')); ?>" class="btn btn-success">
                <i class="voyager-plus"></i> Add New
            </a>
        <?php endif; ?>
    </h1>
    <?php echo $__env->make('voyager::multilingual.language-selector', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>
    <div class="page-content container-fluid">
        <?php echo $__env->make('voyager::alerts', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <div class="row">
            <div class="col-md-12">
                <div class="panel panel-bordered">
                    <div class="panel-body">
                        <table id="dataTable" class="table table-hover">
                            <thead>
                                <tr>
                                    <?php $__currentLoopData = $dataType->browseRows; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $row): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                    <th><?php echo e($row->display_name); ?></th>
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
                                            <?php if(is_field_translatable($data, $row)): ?>
                                                <?php echo $__env->make('voyager::multilingual.input-hidden', [
                                                    '_field_name'  => $row->field,
                                                    '_field_trans' => get_field_translations($data, $row->field)
                                                ], array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                            <?php endif; ?>
                                            <span><?php echo e($data->{$row->field}); ?></span>
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
                                        <?php if(Voyager::can('read_'.$dataType->name)): ?>
                                            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.show', $data->id)); ?>" class="btn-sm btn-warning pull-right">
                                                <i class="voyager-eye"></i> View
                                            </a>
                                        <?php endif; ?>
                                    </td>
                                </tr>
                                <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                            </tbody>
                        </table>
                        <?php if(isset($dataType->server_side) && $dataType->server_side): ?>
                            <div class="pull-left">
                                <div role="status" class="show-res" aria-live="polite">Showing <?php echo e($dataTypeContent->firstItem()); ?> to <?php echo e($dataTypeContent->lastItem()); ?> of <?php echo e($dataTypeContent->total()); ?> entries</div>
                            </div>
                            <div class="pull-right">
                                <?php echo e($dataTypeContent->links()); ?>

                            </div>
                        <?php endif; ?>
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
                    <form action="<?php echo e(route('voyager.'.$dataType->slug.'.destroy', ['id' => '__id'])); ?>" id="delete_form" method="POST">
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
    
    <script>
        $(document).ready(function () {
            <?php if(!$dataType->server_side): ?>
                $('#dataTable').DataTable({ "order": [] });
            <?php endif; ?>
            <?php if($isModelTranslatable): ?>
                $('.side-body').multilingual();
            <?php endif; ?>
        });

        $('td').on('click', '.delete', function(e) {
            $('#delete_form')[0].action = $('#delete_form')[0].action.replace('__id', $(e.target).data('id'));
            $('#delete_modal').modal('show');
        });
    </script>
    <?php if($isModelTranslatable): ?>
        <script src="<?php echo e(voyager_asset('js/multilingual.js')); ?>"></script>
    <?php endif; ?>
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>