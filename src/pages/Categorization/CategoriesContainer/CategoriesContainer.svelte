<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar'
    import {
        Container,
        maxHeight,
        SideMainPadding,
        ContainerWithoutData
    } from "./style";
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
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";

    const addNewCategoryHandler = () => {
        alert("add")
    }

    export let categories: Category[] = [];
    export let status: LoadingStatus = LoadingStatus.None;
</script>

<CommonTitle title="Categories" withButton={true}>
    <Button title="Add new category" onClickHandler={addNewCategoryHandler}/>
</CommonTitle>
<div class="{SideMainPadding}">
    <div class="{Container} {(!categories || categories.length === 0) && ContainerWithoutData}">
        {#if status === LoadingStatus.Loading}
            <LoadingSpinner/>
        {:else if status === LoadingStatus.Finished}
            {#if !categories || categories.length === 0}
                <NoDataFoundBlock/>
            {:else}
                <SimpleBar style="max-height: {maxHeight}px; width: 100%">
                    {#each categories as category (`${category.id}-${category.name}`)}
                        <CategoryRow category={category}/>
                    {/each}
                </SimpleBar>
            {/if}
        {:else if status === LoadingStatus.Error}
            Error fetching data
        {/if}
    </div>
</div>
