<script lang="ts">
    import {_} from 'svelte-i18n'
    import SimpleBar from '@woden/svelte-simplebar'
    import {style} from "./style";
    import {Category} from "../types";
    import CategoryRow from "../CategoryRow/CategoryRow.svelte";
    import ButtonIconNew
        from "../../../common/components/Buttons/ButtonIconNew/ButtonIconNew.svelte";
    import {LoadingStatus} from "../../../stores/utils";
    import LoadingSpinner
        from "../../../common/components/ElementsAndBlocks/LoadingSpinner/LoadingSpinner.svelte";
    import NoDataFoundBlock
        from "../../../common/components/ElementsAndBlocks/NoDataFoundBlock/NoDataFoundBlock.svelte";
    import CommonTitle
        from "../../../common/components/Tables/CommonTitle/CommonTitle.svelte";
    import EditCategoryActionElement
        from "../EditCategoryActionElement/EditCategoryActionElement.svelte";

    export let categories: Category[] = [];
    export let status: LoadingStatus = LoadingStatus.None;
</script>

<CommonTitle title={$_("categories.titles.categories")} withButton={true}>
    <EditCategoryActionElement withButton={true} {status}
                               buttonTitle={$_('categories.buttons.new')}/>
</CommonTitle>
<div class="{style.SideMainPadding}">
    <div class="{style.Container} {(!categories || categories.length === 0) && style.ContainerWithoutData}">
        {#if status === LoadingStatus.Loading}
            <LoadingSpinner/>
        {:else if status === LoadingStatus.Finished}
            {#if !categories || categories.length === 0}
                <NoDataFoundBlock/>
            {:else}
                <SimpleBar style="max-height: {style.maxHeight}px; width: 100%">
                    {#each categories as category (`${category.id}-${category.name}`)}
                        <CategoryRow category={category}/>
                    {/each}
                </SimpleBar>
            {/if}
        {:else if status === LoadingStatus.Error}
            {$_("common.messages.errors.fetching")}
        {/if}
    </div>
</div>
