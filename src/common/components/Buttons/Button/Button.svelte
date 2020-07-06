<script lang="typescript">
    import {
        Button,
        ButtonReset,
        ButtonDefault,
        ButtonPrimary,
        ButtonSecondary,
        Font316Black
    } from './style';

    export let onClickHandler = () => {
    };
    export let title: string = "";
    export let buttonClass: string = "";
    export let secondary: boolean = false;
    export let disabled: boolean = false;
</script>

<button {disabled} class="{ButtonReset} {Button} {ButtonDefault} {secondary ? ButtonSecondary : ButtonPrimary} {Font316Black} {buttonClass}"
        on:click={onClickHandler}>
    {title}
</button>