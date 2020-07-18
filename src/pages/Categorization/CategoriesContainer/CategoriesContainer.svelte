<script lang="ts">
    import {Container} from "./style";
    import {Category} from "../types";
    import CategoryRow from "../CategoryRow/CategoryRow.svelte";

    export let categories: Category[] = [];
</script>


<div class="{Container}">
    {#each categories as category}
        <CategoryRow category={category}/>
    {/each}
</div>