<script lang="ts">
    import {
        Font312Black,
        FlexHorCenter,
        SideMinorPadding,
        ArrowButton,
        MarginedName,
        ArrowButtonExpanded,
        CategoryLine
    } from "./style";
    import {Category} from "../types";
    import {categoryTreeLeftShift} from "./utils";
    import ButtonIconExpandArrow
        from "../../../common/components/Buttons/ButtonIconExpandArrow/ButtonIconExpandArrow.svelte";
    import SingleCategory from "./CategoryRow.svelte";

    const onExpandClick = () => {
        isOpened = !isOpened;
    }

    let isOpened = false;

    export let category: Category = {};
    export let treeLevel: number = 0;
</script>

<div class="{FlexHorCenter} {CategoryLine}"
     style="padding-left: {8 + treeLevel * categoryTreeLeftShift}px">
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
{#if isOpened}
    {#each category.categories as innerCategory}
        <SingleCategory category={innerCategory} treeLevel={treeLevel + 1}/>
    {/each}
{/if}