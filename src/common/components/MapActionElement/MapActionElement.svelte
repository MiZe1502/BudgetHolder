<script lang="typescript">
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

    export let data: MapItemData = [];
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={data[0].name || "Map"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            {#if isDomReady}
                <MapContainer data={[{address: "Moscow, 2-ya Vladimirskaya, 20", name: "test1"}, {address: "Moscow, tverskaya 2", name: "test2"}, {address: "Moscow, Molostovikh, 13", name: "text3"}]} name={data.name}/>
            {/if}
        </PopupContainer>
    {/if}
</div>