<script lang="typescript">
    export let title = "";
    export let dataLength = 0;
</script>

<style>
    .Title {
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-bottom: 8px;
    }

    .TitleText {
        margin-right: 8px;
    }
</style>

<div class="GlobalPadding Title">
    <span class="TitleText Font732Black">{title}</span>
    {#if dataLength}
        <span class="TitleText Font732DarkGray">({dataLength})</span>
    {/if}
</div>