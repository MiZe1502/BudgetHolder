<script lang="typescript">

    import { yandexMapsReady, googleMapsReady } from '../../../stores/maps';

    import { onMount } from 'svelte';

    onMount(async () => {
		// const map = new google.maps.Map(container, {
        //     zoomGmaps,
		// 	centerGmaps,
		// });

        var myMap = new ymaps.Map(container, {
            center: centerYmaps,
            zoom: zoomYmaps
        });

        const points = myMap.geoObjects;
        locations.forEach(location => {
            points.add(new ymaps.Placemark(
                [location.latitude, location.longitude],
                    {balloonContent: location.name},
                    {
            }));
        });
    })

    let container;

    export let locations = [{
        latitude: 55.75361503443606,
        longitude: 37.620883000000006,
        name: 'Test point'
    }];
    export let centerGmaps = {lat: -34.397, lng: 150.644};
    export let centerYmaps = [55.75361503443606, 37.620883000000006];
    export let zoomYmaps = 17;
    export let zoomGmaps = 8;
</script>

<div class="containerWrapper" style="width: 100%; height: 100%; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px;">
    {#if yandexMapsReady && googleMapsReady}
        <div style="width:100%; height: 100%;" bind:this={container}>
        </div>
    {/if}
</div>


