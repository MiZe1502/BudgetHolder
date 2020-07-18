<script lang="typescript">
    import {onMount} from "svelte";
    import {mockData} from "./data";
    import CategoryRow from "./CategoryRow/CategoryRow.svelte";

    onMount(() => {

    })

</script>

<div>
    {#each mockData as category}
        <CategoryRow category={category}/>
    {/each}
</div>