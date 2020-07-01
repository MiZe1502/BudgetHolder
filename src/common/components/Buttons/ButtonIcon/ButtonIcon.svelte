<script lang="typescript">
	import { Button, ButtonReset, ButtonDefault, defaultButtonIconSize } from './style';

    export let onClickHandler = () => {};
    export let width: number = defaultButtonIconSize;
    export let height: number = defaultButtonIconSize;
    export let className: string = "";
</script>

<button class="{Button} {ButtonReset} {ButtonDefault} {className}" on:click={onClickHandler}>
    <slot width={width} height={width} />
</button>