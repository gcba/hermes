<div class="side-menu sidebar-inverse">
    <nav class="navbar navbar-default" role="navigation">
        <div class="side-menu-container">
            <div class="navbar-header">
                <a class="navbar-brand" href="<?php echo e(route('voyager.dashboard')); ?>">
                    <div class="logo-icon-container">
                        <?php $admin_logo_img = Voyager::setting('admin_icon_image', ''); ?>
                        <?php if($admin_logo_img == ''): ?>
                            <img src="<?php echo e(voyager_asset('images/logo-icon-light.png')); ?>" alt="Logo Icon">
                        <?php else: ?>
                            <img src="<?php echo e(Voyager::image($admin_logo_img)); ?>" alt="Logo Icon">
                        <?php endif; ?>
                    </div>
                    <div class="title"><?php echo e(Voyager::setting('admin_title', 'VOYAGER')); ?></div>
                </a>
            </div><!-- .navbar-header -->

            <div class="panel widget center bgimage"
                 style="background-image:url(<?php echo e(Voyager::image( Voyager::setting('admin_bg_image'), config('voyager.assets_path') . '/images/bg.jpg' )); ?>);">
                <div class="dimmer"></div>
                <div class="panel-content">
                    <img src="<?php echo e($user_avatar); ?>" class="avatar" alt="<?php echo e(Auth::user()->name); ?> avatar">
                    <h4><?php echo e(ucwords(Auth::user()->name)); ?></h4>
                    <p><?php echo e(Auth::user()->email); ?></p>

                    <a href="<?php echo e(route('voyager.profile')); ?>" class="btn btn-primary">Profile</a>
                    <div style="clear:both"></div>
                </div>
            </div>

        </div>

        <?php echo menu('admin', 'admin_menu'); ?>

    </nav>
</div>
