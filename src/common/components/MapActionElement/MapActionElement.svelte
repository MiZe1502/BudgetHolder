<script lang="typescript">
    import ButtonIconMap from "../Buttons/ButtonIconMap/ButtonIconMap.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"

    import { mapReady } from '../../../stores/common';

    import { onMount } from 'svelte';

    import { MapActionElementData } from "./utils";
 
    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
        setTimeout(() => {
            map = new google.maps.Map(container, {
                zoom: zoom,
                center: center
            });
        }, 100)

    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    let container;
	let map;
    let zoom = 8;
    let center = {lat: -34.397, lng: 150.644};

    onMount(async () => {
        // map = new google.maps.Map(container, {
        //     zoom: zoom,
        3
        //     center: center
        // });
        console.log(map)
    });

    export let data: MapActionElementData = {};
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer title={data.name || "Map"} isPopupOpened={isPopupOpened} onCloseHandler={onCloseHandler}>
            {#if mapReady}
                <div style="width: 100%; height: 100%; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px;" bind:this={container}>
                    <!-- map is rendered here -->
                </div>
            {/if}
        </PopupContainer>
    {/if}
</div>