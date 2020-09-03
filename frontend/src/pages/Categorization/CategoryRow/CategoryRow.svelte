<script lang="ts">
    import {style} from "./style";
    import {Category} from "../types";
    import {categoryTreeLeftShift} from "./utils";
    import ButtonIconExpandArrow
        from "../../../common/components/Buttons/ButtonIconExpandArrow/ButtonIconExpandArrow.svelte";
    import CategoryRow from "./CategoryRow.svelte";
    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import CategoryActionsElement
        from "../CategoryActionsElement/CategoryActionsElement.svelte";

    const onExpandClick = () => {
        isOpened = !isOpened;
    }

    let isOpened = false;

    export let category: Category = {};
    export let treeLevel: number = 0;
</script>

<div class="{style.FlexHorCenter} {style.CategoryLine} {isOpened && style.CategoryLineExpanded}"
     style="padding-left: {8 + treeLevel * categoryTreeLeftShift}px">
    <div class="{style.FlexHorCenter}">
        {#if category.categories}
            <ButtonIconExpandArrow
                    height={16}
                    width={16}
                    className={isOpened ? style.ArrowButtonExpanded : style.ArrowButton}
                    onClickHandler={onExpandClick}/>
        {/if}
        <div class="{style.Font312Black} {!category.categories && style.MarginedName}">
            {category.name}
        </div>
    </div>
    <CategoryActionsElement data={category}/>
</div>
{#if isOpened}
    {#each category.categories as innerCategory (`${innerCategory.id}-${innerCategory.name}`)}
        <CategoryRow category={innerCategory} treeLevel={treeLevel + 1}/>
    {/each}
{/if}