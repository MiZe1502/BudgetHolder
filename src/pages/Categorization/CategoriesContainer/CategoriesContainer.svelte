<script lang="ts">
    import SimpleBar from '@woden/svelte-simplebar'
    import {Container, maxHeight} from "./style";
    import {Category} from "../types";
    import CategoryRow from "../CategoryRow/CategoryRow.svelte";

    export let categories: Category[] = [];
</script>

<div class="{Container}">
    <SimpleBar style="max-height: {maxHeight}px; width: 100%">
        {#each categories as category}
            <CategoryRow category={category}/>
        {/each}
    </SimpleBar>
</div>