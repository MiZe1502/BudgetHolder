<script lang="ts">
    import {LoadingStatus} from "../../../../stores/utils";
    import {WrapperWhileLoading, FlexHorCenter} from "./style";
    import LoadingSpinner
        from "../../ElementsAndBlocks/LoadingSpinner/LoadingSpinner.svelte";
    import Button from "../Button/Button.svelte";

    export let status: LoadingStatus = LoadingStatus.None;
    export let onClickHandler = () => {
    };
    export let title: string = "";
    export let buttonClass: string = "";
    export let secondary: boolean = false;
    export let disabled: boolean = false;
</script>


<div class="{FlexHorCenter} {status === LoadingStatus.Loading && WrapperWhileLoading}">
    {#if status === LoadingStatus.Loading}
        <LoadingSpinner width={30} height={30}/>
    {:else if status === LoadingStatus.Finished}
        <Button {title} {buttonClass} {secondary} {disabled} {onClickHandler}/>
    {/if}
</div>