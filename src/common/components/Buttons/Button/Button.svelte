<script lang="typescript">
	import { Button, ButtonReset, ButtonDefault, Font316Black } from './style';

    export let onClickHandler = () => {};
    export let title: string = "";
    export let buttonClass: string = "";
</script>

<button class="{ButtonReset} {Button} {ButtonDefault} {Font316Black} {buttonClass}" on:click={onClickHandler}>
    {title}
</button>