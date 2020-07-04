<script lang="typescript">
	import { Button, ButtonReset, ButtonDefault, defaultButtonIconSize } from './style';

    export let onClickHandler = () => {};
    export let width: number = defaultButtonIconSize;
    export let height: number = defaultButtonIconSize;
    export let className: string = "";
    export let disabled: boolean = false;
</script>

<button disabled={disabled} class="{Button} {ButtonReset} {ButtonDefault} {className}" on:click={onClickHandler}>
    <slot width={width} height={width} />
</button>