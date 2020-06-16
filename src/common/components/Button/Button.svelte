<script lang="typescript">
	import { Button, Font316Black } from './style';

    export let onClickHandler = () => {};
    export let title: string = "";
</script>

<button class="{Button} {Font316Black}" onClick={onClickHandler}>
    {title}
</button>