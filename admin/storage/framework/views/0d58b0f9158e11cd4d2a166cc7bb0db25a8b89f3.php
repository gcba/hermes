<?php if(Auth::check()): ?>
    <!-- Left side column. contains the sidebar -->
    <aside class="main-sidebar">
      <!-- sidebar: style can be found in sidebar.less -->
      <section class="sidebar">
        <!-- Sidebar user panel -->
        <div class="user-panel">
          <div class="pull-left image">
            <img src="https://placehold.it/160x160/00a65a/ffffff/&text=<?php echo e(mb_substr(Auth::user()->name, 0, 1)); ?>" class="img-circle" alt="User Image">
          </div>
          <div class="pull-left info">
            <p><?php echo e(Auth::user()->name); ?></p>
            <a href="#"><i class="fa fa-circle text-success"></i> Online</a>
          </div>
        </div>
        <!-- sidebar menu: : style can be found in sidebar.less -->
        <ul class="sidebar-menu">
          <li class="header"><?php echo e(trans('backpack::base.administration')); ?></li>
          <!-- ================================================ -->
          <!-- ==== Recommended place for admin menu items ==== -->
          <!-- ================================================ -->
          <li><a href="<?php echo e(url(config('backpack.base.route_prefix', 'admin').'/dashboard')); ?>"><i class="fa fa-dashboard"></i> <span><?php echo e(trans('backpack::base.dashboard')); ?></span></a></li>


          <!-- ======================================= -->
          <li class="header"><?php echo e(trans('backpack::base.user')); ?></li>
          <li><a href="<?php echo e(url(config('backpack.base.route_prefix', 'admin').'/logout')); ?>"><i class="fa fa-sign-out"></i> <span><?php echo e(trans('backpack::base.logout')); ?></span></a></li>
        </ul>
      </section>
      <!-- /.sidebar -->
    </aside>
<?php endif; ?>
