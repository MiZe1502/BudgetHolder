<script lang="typescript">
    import { MapContainerWrapper } from "./style";

    import ButtonIconMap from "../Buttons/ButtonIconMap/ButtonIconMap.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"
    import MapContainer from "../MapContainer/MapContainer.svelte";

    import { onMount } from 'svelte';

    import { MapActionElementData, MapItemData } from "./utils";
 
    let isPopupOpened = false;

    let isDomReady = false;

    const onClickHandler = () => {
        isPopupOpened = true;
        setTimeout(() => {
            isDomReady = true;
        }, 100)

    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: MapItemData[] = [];
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={data[0].name || "Map"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            {#if isDomReady}
                <MapContainer data={data} wrapperClass={MapContainerWrapper}/>
            {/if}
        </PopupContainer>
    {/if}
</div>