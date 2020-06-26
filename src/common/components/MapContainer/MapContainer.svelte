<script lang="typescript">

    import { MapItemData } from '../MapActionElement/utils';

    import { yandexMapsReady, googleMapsReady } from '../../../stores/maps';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    import { onMount } from 'svelte';

    onMount(async () => {
        const placemarkBuilder: PlacemarkBuilder = new YandexMapsPlacemarkBuilder();
        const mapBuilder: MapsBuilder = new YandexMapsBuilder(container, placemarkBuilder);

        if (data.length > 1) {
            mapBuilder.findMultipleAddressesAndCreateMap(data);
        } else {
            mapBuilder.findSingleAddressAndCreateMap(data[0]);
        }
    })

    let container;

    export let data: MapItemData[] = [];
</script>

<div class="containerWrapper" style="width: 100%; height: 550px; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px;">
    {#if yandexMapsReady && googleMapsReady}
        <div style="width:100%; height: 100%;" bind:this={container}>
        </div>
    {/if}
</div>


