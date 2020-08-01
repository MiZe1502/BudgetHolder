<script lang="ts">
    import {FlexVert, Delimeter, SingleLine} from "./style";
    import SimpleTextElement
        from "../SimpleTextElement/SimpleTextElement.svelte";

    export let data: string[] = [];
</script>

<div class="{FlexVert}">
    {#if data}
        {#each data as category (category)}
            <SimpleTextElement className={SingleLine} data={category}/>
            <div class="{Delimeter}"/>
        {/each}
    {:else}
        {"\u2014"}
    {/if}
</div>
