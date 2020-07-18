<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar'
    import {Container, maxHeight} from "./style";
    import {Category} from "../types";
    import CategoryRow from "../CategoryRow/CategoryRow.svelte";
    import ButtonIconNew
        from "../../../common/components/Buttons/ButtonIconNew/ButtonIconNew.svelte";

    export let categories: Category[] = [];
</script>

<ButtonIconNew />
<div class="{Container}">
    <SimpleBar style="max-height: {maxHeight}px; width: 100%">
        {#each categories as category (`${category.id}-${category.name}`)}
            <CategoryRow category={category}/>
        {/each}
    </SimpleBar>
</div>