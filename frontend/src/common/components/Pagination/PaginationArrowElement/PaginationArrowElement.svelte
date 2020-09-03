<script lang="typescript">
    import { Direction } from "./utils";

    import ButtonIconRightArrow from "../../Buttons/ButtonIconRightArrow/ButtonIconRightArrow.svelte";
    import ButtonIconLeftArrow from "../../Buttons/ButtonIconLeftArrow/ButtonIconLeftArrow.svelte";

    export let disabled: boolean = false;
    export let direction: Direction;
    export let onClickHandler = () => {};
</script>

<div>
    {#if direction === Direction.Left}
        <ButtonIconLeftArrow disabled={disabled} onClickHandler={onClickHandler}/>
    {/if}
    {#if direction === Direction.Right}
        <ButtonIconRightArrow disabled={disabled} onClickHandler={onClickHandler}/>
    {/if}
</div>