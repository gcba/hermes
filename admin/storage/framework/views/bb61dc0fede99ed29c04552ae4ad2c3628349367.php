<?php $__env->startSection('page_title','All '.$dataType->display_name_plural); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="<?php echo e($dataType->icon); ?>"></i> <?php echo e($dataType->display_name_plural); ?>

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
                    <div class="panel-body table-responsive">
                        <table id="dataTable" class="row table table-hover">
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
                                            <?php $options = json_decode($row->details); ?>
                                            <?php if($row->type == 'image'): ?>
                                                <img src="<?php if( strpos($data->{$row->field}, 'http://') === false && strpos($data->{$row->field}, 'https://') === false): ?><?php echo e(Voyager::image( $data->{$row->field} )); ?><?php else: ?><?php echo e($data->{$row->field}); ?><?php endif; ?>" style="width:100px">
                                            <?php elseif($row->type == 'select_multiple'): ?>
                                                <?php if(property_exists($options, 'relationship')): ?>

                                                    <?php $__currentLoopData = $data->{$row->field}; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $item): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                                        <?php if($item->{$row->field . '_page_slug'}): ?>
                                                        <a href="<?php echo e($item->{$row->field . '_page_slug'}); ?>"><?php echo e($item->{$row->field}); ?></a><?php if(!$loop->last): ?>, <?php endif; ?>
                                                        <?php else: ?>
                                                        <?php echo e($item->{$row->field}); ?>

                                                        <?php endif; ?>
                                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>

                                                    
                                                <?php elseif(property_exists($options, 'options')): ?>
                                                    <?php $__currentLoopData = $data->{$row->field}; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $item): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                                     <?php echo e($options->options->{$item} . (!$loop->last ? ', ' : '')); ?>

                                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                                <?php endif; ?>

                                            <?php elseif($row->type == 'select_dropdown' && property_exists($options, 'options')): ?>

                                                <?php if($data->{$row->field . '_page_slug'}): ?>
                                                    <a href="<?php echo e($data->{$row->field . '_page_slug'}); ?>"><?php echo $options->options->{$data->{$row->field}}; ?></a>
                                                <?php else: ?>
                                                    <?php echo $options->options->{$data->{$row->field}}; ?>

                                                <?php endif; ?>


                                            <?php elseif($row->type == 'select_dropdown' && $data->{$row->field . '_page_slug'}): ?>
                                                <a href="<?php echo e($data->{$row->field . '_page_slug'}); ?>"><?php echo e($data->{$row->field}); ?></a>
                                            <?php elseif($row->type == 'date'): ?>
                                            <?php echo e($options && property_exists($options, 'format') ? \Carbon\Carbon::parse($data->{$row->field})->formatLocalized($options->format) : $data->{$row->field}); ?>

                                            <?php elseif($row->type == 'checkbox'): ?>
                                                <?php if($options && property_exists($options, 'on') && property_exists($options, 'off')): ?>
                                                    <?php if($data->{$row->field}): ?>
                                                    <span class="label label-info"><?php echo e($options->on); ?></span>
                                                    <?php else: ?>
                                                    <span class="label label-primary"><?php echo e($options->off); ?></span>
                                                    <?php endif; ?>
                                                <?php else: ?>
                                                <?php echo e($data->{$row->field}); ?>

                                                <?php endif; ?>
                                            <?php elseif($row->type == 'text'): ?>
                                                <?php echo $__env->make('voyager::multilingual.input-hidden-bread-browse', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                                <div class="readmore"><?php echo e(strlen( $data->{$row->field} ) > 200 ? substr($data->{$row->field}, 0, 200) . ' ...' : $data->{$row->field}); ?></div>
                                            <?php elseif($row->type == 'text_area'): ?>
                                                <?php echo $__env->make('voyager::multilingual.input-hidden-bread-browse', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                                <div class="readmore"><?php echo e(strlen( $data->{$row->field} ) > 200 ? substr($data->{$row->field}, 0, 200) . ' ...' : $data->{$row->field}); ?></div>
                                            <?php elseif($row->type == 'rich_text_box'): ?>
                                                <?php echo $__env->make('voyager::multilingual.input-hidden-bread-browse', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                                <div class="readmore"><?php echo e(strlen( strip_tags($data->{$row->field}, '<b><i><u>') ) > 200 ? substr(strip_tags($data->{$row->field}, '<b><i><u>'), 0, 200) . ' ...' : strip_tags($data->{$row->field}, '<b><i><u>')); ?></div>
                                            <?php else: ?>
                                                <?php echo $__env->make('voyager::multilingual.input-hidden-bread-browse', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                                <span><?php echo e($data->{$row->field}); ?></span>
                                            <?php endif; ?>
                                        </td>
                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                    <td class="no-sort no-click" id="bread-actions">
                                        <?php if(Voyager::can('delete_'.$dataType->name)): ?>
                                            <a href="javascript:;" title="Delete" class="btn btn-sm btn-danger pull-right delete" data-id="<?php echo e($data->id); ?>" id="delete-<?php echo e($data->id); ?>">
                                                <i class="voyager-trash"></i> <span class="hidden-xs hidden-sm">Delete</span>
                                            </a>
                                        <?php endif; ?>
                                        <?php if(Voyager::can('edit_'.$dataType->name)): ?>
                                            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.edit', $data->id)); ?>" title="Edit" class="btn btn-sm btn-primary pull-right edit">
                                                <i class="voyager-edit"></i> <span class="hidden-xs hidden-sm">Edit</span>
                                            </a>
                                        <?php endif; ?>
                                        <?php if(Voyager::can('read_'.$dataType->name)): ?>
                                            <a href="<?php echo e(route('voyager.'.$dataType->slug.'.show', $data->id)); ?>" title="View" class="btn btn-sm btn-warning pull-right">
                                                <i class="voyager-eye"></i> <span class="hidden-xs hidden-sm">View</span>
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
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title"><i class="voyager-trash"></i> Are you sure you want to delete
                        this <?php echo e(strtolower($dataType->display_name_singular)); ?>?</h4>
                </div>
                <div class="modal-footer">
                    <form action="<?php echo e(route('voyager.'.$dataType->slug.'.index')); ?>" id="delete_form" method="POST">
                        <?php echo e(method_field("DELETE")); ?>

                        <?php echo e(csrf_field()); ?>

                        <input type="submit" class="btn btn-danger pull-right delete-confirm"
                                 value="Yes, delete this <?php echo e(strtolower($dataType->display_name_singular)); ?>">
                    </form>
                    <button type="button" class="btn btn-default pull-right" data-dismiss="modal">Cancel</button>
                </div>
            </div><!-- /.modal-content -->
        </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->
<?php $__env->stopSection(); ?>

<?php $__env->startSection('css'); ?>
<?php if(!$dataType->server_side && config('dashboard.data_tables.responsive')): ?>
<link rel="stylesheet" href="<?php echo e(voyager_asset('lib/css/responsive.dataTables.min.css')); ?>">
<?php endif; ?>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>
    <!-- DataTables -->
    <?php if(!$dataType->server_side && config('dashboard.data_tables.responsive')): ?>
        <script src="<?php echo e(voyager_asset('lib/js/dataTables.responsive.min.js')); ?>"></script>
    <?php endif; ?>
    <?php if($isModelTranslatable): ?>
        <script src="<?php echo e(voyager_asset('js/multilingual.js')); ?>"></script>
    <?php endif; ?>
    <script>
        $(document).ready(function () {
            <?php if(!$dataType->server_side): ?>
                var table = $('#dataTable').DataTable({
                    "order": []
                    <?php if(config('dashboard.data_tables.responsive')): ?>, responsive: true <?php endif; ?>
                });
            <?php endif; ?>

            <?php if($isModelTranslatable): ?>
                $('.side-body').multilingual();
            <?php endif; ?>
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
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>