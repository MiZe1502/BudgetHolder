<script lang="typescript">
	import { Button, ButtonReset, ButtonDefault, Font316Black } from './style';

    export let onClickHandler = () => {};
    export let title: string = "";
</script>

<button class="{ButtonReset} {Button} {ButtonDefault} {Font316Black}" on:click={onClickHandler}>
    {title}
</button>