<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {style} from "./style";

    import Button from '../../Buttons/Button/Button.svelte'

    export let title = "";
    export let dataLength = 0;

    export let withButton: boolean = false;
    export let buttonTitle: string = "";
    export let buttonClickHandler = () => {
    };
</script>

<div class="{style.SideMainPadding} {style.Title} {style.FlexHorCenter}">
    <span class="{style.TitleText} {style.Font732Black}">{$_(title)}</span>
    {#if dataLength}
        <span class="{style.TitleText} {style.Font732DarkGray}">({dataLength})</span>
    {/if}
    {#if withButton}
        <slot/>
    {/if}
</div>