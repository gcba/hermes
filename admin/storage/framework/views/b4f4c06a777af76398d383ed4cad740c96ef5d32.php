<?php if(config('voyager.show_dev_tips')): ?>
    <div class="container-fluid">
        <div class="alert alert-info">
            <strong>How To Use:</strong>
            <p>You can output <?php echo e(!empty($menu) ? 'this' : 'a'); ?> menu anywhere on your site by calling <code>menu('<?php echo e(!empty($menu) ? $menu->name : 'name'); ?>')</code></p>
        </div>
    </div>
<?php endif; ?>
