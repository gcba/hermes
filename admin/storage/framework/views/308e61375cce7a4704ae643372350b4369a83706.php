<?php $__env->startSection('page_title','View '.$dataType->display_name_singular); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="<?php echo e($dataType->icon); ?>"></i> Viewing <?php echo e(ucfirst($dataType->display_name_singular)); ?> &nbsp;

        <?php if(Voyager::can('edit_'.$dataType->name)): ?>
        <a href="<?php echo e(route('voyager.'.$dataType->slug.'.edit', $dataTypeContent->getKey())); ?>" class="btn btn-info">
            <span class="glyphicon glyphicon-pencil"></span>&nbsp;
            Edit
        </a>
        <?php endif; ?>
        <a href="<?php echo e(route('voyager.'.$dataType->slug.'.index')); ?>" class="btn btn-warning">
            <span class="glyphicon glyphicon-list"></span>&nbsp;
            Return to List
        </a>        
    </h1>
    <?php echo $__env->make('voyager::multilingual.language-selector', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>
    <div class="page-content container-fluid">
        <div class="row">
            <div class="col-md-12">

                <div class="panel panel-bordered" style="padding-bottom:5px;">

                    <!-- /.box-header -->
                    <!-- form start -->

                    <?php $__currentLoopData = $dataType->readRows; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $row): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                        <?php  $rowDetails = json_decode($row->details);  ?>

                        <div class="panel-heading" style="border-bottom:0;">
                            <h3 class="panel-title"><?php echo e($row->display_name); ?></h3>
                        </div>

                        <div class="panel-body" style="padding-top:0;">
                            <?php if($row->type == "image"): ?>
                                <img class="img-responsive"
                                     src="<?php echo e(Voyager::image($dataTypeContent->{$row->field})); ?>">
                            <?php elseif($row->type == 'select_dropdown' && property_exists($rowDetails, 'options') &&
                                    !empty($rowDetails->options->{$dataTypeContent->{$row->field}})
                            ): ?>

                                <?php echo $rowDetails->options->{$dataTypeContent->{$row->field}};?>
                            <?php elseif($row->type == 'select_dropdown' && $dataTypeContent->{$row->field . '_page_slug'}): ?>
                                <a href="<?php echo e($dataTypeContent->{$row->field . '_page_slug'}); ?>"><?php echo e($dataTypeContent->{$row->field}); ?></a>
                            <?php elseif($row->type == 'select_multiple'): ?>
                                <?php if(property_exists($rowDetails, 'relationship')): ?>

                                    <?php $__currentLoopData = $dataTypeContent->{$row->field}; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $item): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                        <?php if($item->{$row->field . '_page_slug'}): ?>
                                        <a href="<?php echo e($item->{$row->field . '_page_slug'}); ?>"><?php echo e($item->{$row->field}); ?></a><?php if(!$loop->last): ?>, <?php endif; ?>
                                        <?php else: ?>
                                        <?php echo e($item->{$row->field}); ?>

                                        <?php endif; ?>
                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>

                                <?php elseif(property_exists($rowDetails, 'options')): ?>
                                    <?php $__currentLoopData = $dataTypeContent->{$row->field}; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $item): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                     <?php echo e($rowDetails->options->{$item} . (!$loop->last ? ', ' : '')); ?>

                                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                <?php endif; ?>
                            <?php elseif($row->type == 'date'): ?>
                                <?php echo e($rowDetails && property_exists($rowDetails, 'format') ? \Carbon\Carbon::parse($dataTypeContent->{$row->field})->formatLocalized($rowDetails->format) : $dataTypeContent->{$row->field}); ?>

                            <?php elseif($row->type == 'checkbox'): ?>
                                <?php if($rowDetails && property_exists($rowDetails, 'on') && property_exists($rowDetails, 'off')): ?>
                                    <?php if($dataTypeContent->{$row->field}): ?>
                                    <span class="label label-info"><?php echo e($rowDetails->on); ?></span>
                                    <?php else: ?>
                                    <span class="label label-primary"><?php echo e($rowDetails->off); ?></span>
                                    <?php endif; ?>
                                <?php else: ?>
                                <?php echo e($dataTypeContent->{$row->field}); ?>

                                <?php endif; ?>
                            <?php elseif($row->type == 'rich_text_box'): ?>
                                <?php echo $__env->make('voyager::multilingual.input-hidden-bread-read', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                <p><?php echo e(strip_tags($dataTypeContent->{$row->field}, '<b><i><u>')); ?></p>
                            <?php else: ?>
                                <?php echo $__env->make('voyager::multilingual.input-hidden-bread-read', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
                                <p><?php echo e($dataTypeContent->{$row->field}); ?></p>
                            <?php endif; ?>
                        </div><!-- panel-body -->
                        <?php if(!$loop->last): ?>
                            <hr style="margin:0;">
                        <?php endif; ?>
                    <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>

                </div>
            </div>
        </div>
    </div>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>
    <?php if($isModelTranslatable): ?>
    <script>
        $(document).ready(function () {
            $('.side-body').multilingual();
        });
    </script>
    <script src="<?php echo e(voyager_asset('js/multilingual.js')); ?>"></script>
    <?php endif; ?>
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>