<script lang="ts">
    import {
        Font312Black,
        FlexHorCenter,
        SideMinorPadding,
        ArrowButton,
        MarginedName,
        ArrowButtonExpanded,
        CategoryLineExpanded,
        CategoryLine
    } from "./style";
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

<div class="{FlexHorCenter} {CategoryLine} {isOpened && CategoryLineExpanded}"
     style="padding-left: {8 + treeLevel * categoryTreeLeftShift}px">
    <div class="{FlexHorCenter}">
        {#if category.categories}
            <ButtonIconExpandArrow
                    height={16}
                    width={16}
                    className={isOpened ? ArrowButtonExpanded : ArrowButton}
                    onClickHandler={onExpandClick}/>
        {/if}
        <div class="{Font312Black} {!category.categories && MarginedName}">
            {category.name}
        </div>
    </div>
    <CategoryActionsElement data={category}/>
<!--    <div class="{FlexHorCenter}">-->
<!--        <ButtonIconEdit width={16} height={16}/>-->
<!--        <ButtonIconRemove width={16} height={16}/>-->
<!--    </div>-->
</div>
{#if isOpened}
    {#each category.categories as innerCategory (`${innerCategory.id}-${innerCategory.name}`)}
        <CategoryRow category={innerCategory} treeLevel={treeLevel + 1}/>
    {/each}
{/if}