<script lang="typescript">
    import ButtonIconMap from "../Buttons/ButtonIconMap/ButtonIconMap.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"

    import { MapActionElementData } from "./utils";
 
    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: MapActionElementData = {};
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={data.name || "Map"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            Popup with map
        </PopupContainer>
    {/if}
</div>