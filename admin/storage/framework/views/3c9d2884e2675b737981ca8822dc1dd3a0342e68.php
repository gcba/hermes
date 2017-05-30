<?php $__env->startSection('css'); ?>
    <meta name="csrf-token" content="<?php echo e(csrf_token()); ?>">
<?php $__env->stopSection(); ?>

<?php $__env->startSection('page_header'); ?>
    <h1 class="page-title">
        <i class="<?php echo e($dataType->icon); ?>"></i> <?php if(isset($dataTypeContent->id)): ?><?php echo e('Edit'); ?><?php else: ?><?php echo e('New'); ?><?php endif; ?> <?php echo e($dataType->display_name_singular); ?>

    </h1>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('content'); ?>
    <div class="page-content container-fluid">
        <?php echo $__env->make('voyager::alerts', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
        <div class="row">
            <div class="col-md-12">

                <div class="panel panel-bordered">

                    <div class="panel-heading">
                        <h3 class="panel-title"><?php if(isset($dataTypeContent->id)): ?><?php echo e('Edit'); ?><?php else: ?><?php echo e('Add New'); ?><?php endif; ?> <?php echo e($dataType->display_name_singular); ?></h3>
                    </div>
                    <!-- /.box-header -->
                    <!-- form start -->
                    <form class="form-edit-add" role="form"
                          action="<?php if(isset($dataTypeContent->id)): ?><?php echo e(route('voyager.'.$dataType->slug.'.update', $dataTypeContent->id)); ?><?php else: ?><?php echo e(route('voyager.'.$dataType->slug.'.store')); ?><?php endif; ?>"
                          method="POST" enctype="multipart/form-data">

                        <!-- PUT Method if we are editing -->
                        <?php if(isset($dataTypeContent->id)): ?>
                            <?php echo e(method_field("PUT")); ?>

                        <?php endif; ?>

                        <!-- CSRF TOKEN -->
                        <?php echo e(csrf_field()); ?>


                        <div class="panel-body">

                            <?php if(count($errors) > 0): ?>
                                <div class="alert alert-danger">
                                    <ul>
                                        <?php $__currentLoopData = $errors->all(); $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $error): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                            <li><?php echo e($error); ?></li>
                                        <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                    </ul>
                                </div>
                            <?php endif; ?>

                            <?php $__currentLoopData = $dataType->addRows; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $row): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                <div class="form-group">
                                    <label for="name"><?php echo e($row->display_name); ?></label>

                                    <?php echo Voyager::formField($row, $dataType, $dataTypeContent); ?>


                                </div>
                            <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>

                            <label for="permission">Permissions</label><br>
                            <a href="#" class="permission-select-all">Select All</a> / <a href="#"  class="permission-deselect-all">Deselect All</a>
                            <ul class="permissions checkbox">
                                <?php
                                    $role_permissions = (isset($dataTypeContent)) ? $dataTypeContent->permissions->pluck('key')->toArray() : [];
                                ?>
                                <?php $__currentLoopData = TCG\Voyager\Models\Permission::all()->groupBy('table_name')->sortBy('table_name'); $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $table => $permission): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                    <li>
                                        <input type="checkbox" id="<?php echo e($table); ?>" class="permission-group">
                                        <label for="<?php echo e($table); ?>"><strong><?php echo e(ucwords($table)); ?></strong></label>
                                        <ul>
                                            <?php $__currentLoopData = $permission; $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $perm): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
                                                <li>
                                                    <input type="checkbox" id="permission-<?php echo e($perm->id); ?>" name="permissions[]" class="the-permission" value="<?php echo e($perm->id); ?>" <?php if(in_array($perm->key, $role_permissions)): ?> checked <?php endif; ?>>
                                                    <label for="permission-<?php echo e($perm->id); ?>"><?php echo e(title_case(str_replace('_', ' ', $perm->key))); ?></label>
                                                </li>
                                            <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                                        </ul>
                                    </li>
                                <?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>
                            </ul>
                        </div><!-- panel-body -->
                        <div class="panel-footer">
                            <button type="submit" class="btn btn-primary">Submit</button>
                        </div>
                    </form>

                    <iframe id="form_target" name="form_target" style="display:none"></iframe>
                    <form id="my_form" action="<?php echo e(route('voyager.upload')); ?>" target="form_target" method="post"
                          enctype="multipart/form-data" style="width:0;height:0;overflow:hidden">
                        <?php echo e(csrf_field()); ?>

                        <input name="image" id="upload_file" type="file"
                               onchange="$('#my_form').submit();this.value='';">
                        <input type="hidden" name="type_slug" id="type_slug" value="<?php echo e($dataType->slug); ?>">
                    </form>

                </div>
            </div>
        </div>
    </div>
<?php $__env->stopSection(); ?>

<?php $__env->startSection('javascript'); ?>
    <script>
        $('document').ready(function () {
            $('.toggleswitch').bootstrapToggle();

            $('.permission-group').on('change', function(){
                $(this).siblings('ul').find("input[type='checkbox']").prop('checked', this.checked);
            });

            $('.permission-select-all').on('click', function(){
                $('ul.permissions').find("input[type='checkbox']").prop('checked', true);
                return false;
            });

            $('.permission-deselect-all').on('click', function(){
                $('ul.permissions').find("input[type='checkbox']").prop('checked', false);
                return false;
            });

            function parentChecked(){
                $('.permission-group').each(function(){
                    var allChecked = true;
                    $(this).siblings('ul').find("input[type='checkbox']").each(function(){
                        if(!this.checked) allChecked = false;
                    });
                    $(this).prop('checked', allChecked);
                });
            }

            parentChecked();

            $('.the-permission').on('change', function(){
                parentChecked();
            });
        });
    </script>
    <script src="<?php echo e(voyager_asset('lib/js/tinymce/tinymce.min.js')); ?>"></script>
    <script src="<?php echo e(voyager_asset('js/voyager_tinymce.js')); ?>"></script>
<?php $__env->stopSection(); ?>

<?php echo $__env->make('voyager::master', array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>