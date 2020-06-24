<script lang="typescript">

    import { yandexMapsReady, googleMapsReady } from '../../../stores/maps';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    import { onMount } from 'svelte';

    onMount(async () => {
		// const map = new google.maps.Map(container, {
        //     zoomGmaps,
		// 	centerGmaps,
		// });



        // const points = myMap.geoObjects;
        // locations.forEach(location => {
        //     points.add(new ymaps.Placemark(
        //         [location.latitude, location.longitude],
        //             {balloonContent: location.name},
        //             {
        //     }));
        // });

        const config = {
            name,
            address,
        }
        const placemarkBuilder: PlacemarkBuilder = new YandexMapsPlacemarkBuilder(config);
        const mapBuilder: MapBuilder = new YandexMapsBuilder(container, placemarkBuilder);

        mapBuilder.findSingleAddressAndCreateMap(config.address);
    })

    let container;

    export let locations = [{
        latitude: 55.75361503443606,
        longitude: 37.620883000000006,
        name: 'Test point'
    }];
    export let centerGmaps = {lat: -34.397, lng: 150.644};
    export let centerYmaps = [55.75361503443606, 37.620883000000006];
    export let zoomGmaps = 8;

    export let address: string = "";
    export let name: string = "";
</script>

<div class="containerWrapper" style="width: 100%; height: 550px; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px;">
    {#if yandexMapsReady && googleMapsReady}
        <div style="width:100%; height: 100%;" bind:this={container}>
        </div>
    {/if}
</div>


