<script lang="typescript">
    import ButtonIconMap from "../Buttons/ButtonIconMap/ButtonIconMap.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"
    import MapContainer from "../MapContainer/MapContainer.svelte";

    import { onMount } from 'svelte';

    import { MapActionElementData } from "./utils";
 
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

    export let data: MapActionElementData = {};
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={data.name || "Map"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            {#if isDomReady}
                <MapContainer address={data.address} name={data.name}/>
            {/if}
        </PopupContainer>
    {/if}
</div>