<script lang="typescript">
    import ButtonIconMap from "../Buttons/ButtonIconMap/ButtonIconMap.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"
 
    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}/>
    {/if}
</div>