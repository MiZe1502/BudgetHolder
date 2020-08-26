<script lang="ts">
    import {style} from "./style";
    import SimpleTextElement
        from "../SimpleTextElement/SimpleTextElement.svelte";

    export let data: string[] = [];
</script>

<div class="{style.FlexVert}">
    {#if data}
        {#each data as category (category)}
            <SimpleTextElement className={style.SingleLine} data={category}/>
            <div class="{style.Delimiter}"/>
        {/each}
    {:else}
        {"\u2014"}
    {/if}
</div>
