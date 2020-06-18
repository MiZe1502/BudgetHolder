<script lang="typescript">
    import { Font732Black, Font732DarkGray, Title, TitleText, SideMainPadding, FlexHorCenter} from "./style";

    import Button from '../Buttons/Button/Button.svelte'

    export let title = "";
    export let dataLength = 0;

    export let withButton: boolean = false;
    export let buttonTitle: string = "";
    export let buttonClickHandler = () => {};
</script>

<div class="{SideMainPadding} {Title} {FlexHorCenter}">
    <span class="{TitleText} {Font732Black}">{title}</span>
    {#if dataLength}
        <span class="{TitleText} {Font732DarkGray}">({dataLength})</span>
    {/if}
    {#if withButton}
        <Button title={buttonTitle} onClickHandler={buttonClickHandler}/>
    {/if}
</div>