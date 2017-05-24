<ul class="nav navbar-nav">

<?php 
    if (Voyager::translatable($items)) {
        $items = $items->load('translations');
    }

 ?>

<?php $__currentLoopData = $items->sortBy('order'); $__env->addLoop($__currentLoopData); foreach($__currentLoopData as $item): $__env->incrementLoopIndices(); $loop = $__env->getLastLoop(); ?>
    
    <?php 
        $originalItem = $item;
        if (Voyager::translatable($item)) {
            $item = $item->translate($options->locale);
        }

        // TODO - still a bit ugly - can move some of this stuff off to a helper in the future.
        $listItemClass = [];
        $styles = null;
        $linkAttributes = null;

        if(url($item->link()) == url()->current())
        {
            array_push($listItemClass,'active');
        }

        // With Children Attributes
        if(!$originalItem->children->isEmpty())
        {
            foreach($originalItem->children as $children)
            {
                if(url($children->link()) == url()->current())
                {
                    array_push($listItemClass,'active');
                }
            }
            $linkAttributes =  'href="#' . str_slug($item->title, '-') .'-dropdown-element" data-toggle="collapse" aria-expanded="'. (in_array('active', $listItemClass) ? 'true' : 'false').'"';
            array_push($listItemClass, 'dropdown');
        }
        else
        {
            $linkAttributes =  'href="' . url($item->link()) .'"';
        }

        // Permission Checker
        $self_prefix = str_replace('/', '\/', $options->user->prefix);
        $slug = str_replace('/', '', preg_replace('/^\/'.$self_prefix.'/', '', $item->link()));

        if ($slug != '') {
            // Get dataType using slug
            $dataType = $options->user->dataTypes->first(function ($value) use ($slug) {
                return $value->slug == $slug;
            });

            if ($dataType) {
                // Check if datatype permission exist
                $exist = $options->user->permissions->first(function ($value) use ($dataType) {
                    return $value->key == 'browse_'.$dataType->name;
                });
            } else {
                // Check if admin permission exists
                $exist = $options->user->permissions->first(function ($value) use ($slug) {
                    return $value->key == 'browse_'.$slug && is_null($value->table_name);
                });
            }

            if ($exist) {
                // Check if current user has access
                if (!in_array($exist->key, $options->user->user_permissions)) {
                    continue;
                }
            }
        }
        
     ?>

    <li class="<?php echo e(implode(" ", $listItemClass)); ?>">
        <a <?php echo $linkAttributes; ?> target="<?php echo e($item->target); ?>">
            <span class="icon <?php echo e($item->icon_class); ?>"></span>
            <span class="title"><?php echo e($item->title); ?></span>
        </a>
        <?php if(!$originalItem->children->isEmpty()): ?>
        <div id="<?php echo e(str_slug($originalItem->title, '-')); ?>-dropdown-element" class="panel-collapse collapse <?php echo e((in_array('active', $listItemClass) ? 'in' : '')); ?>">
            <div class="panel-body">
                <?php echo $__env->make('voyager::menu.admin_menu', ['items' => $originalItem->children, 'options' => $options, 'innerLoop' => true], array_except(get_defined_vars(), array('__data', '__path')))->render(); ?>
            </div>
        </div>
        <?php endif; ?>
    </li>
<?php endforeach; $__env->popLoop(); $loop = $__env->getLastLoop(); ?>

</ul>
