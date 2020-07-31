<script lang="ts">
    import {FlexVert} from "./style";
    import SimpleTextElement
        from "../SimpleTextElement/SimpleTextElement.svelte";

    export let data: string[] = [];
</script>

<div class="{FlexVert}">
    {#if data}
        {#each data as category (category)}
            <SimpleTextElement data={category}/>
        {/each}
    {:else}
        {"\u2014"}
    {/if}
</div>
