<script lang="typescript">
    import Button from '../Button/Button.svelte'

    export let title = "";
    export let dataLength = 0;

    export let withButton: boolean = false;
    export let buttonTitle: string = "";
    export let buttonClickHandler = () => {};
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
    {#if withButton}
        <Button title={buttonTitle} onClickHandler={buttonClickHandler}/>
    {/if}
</div>