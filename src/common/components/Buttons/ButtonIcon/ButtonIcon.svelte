<script lang="typescript">
	import { Button, ButtonReset, ButtonDefault } from './style';

    export let onClickHandler = () => {};
    export let width: number = 24;
    export let height: number = 24;
    export let className: string = "";
</script>

<button class="{Button} {ButtonReset} {ButtonDefault} {className}" on:click={onClickHandler}>
    <slot width={24} height={24} />
</button>