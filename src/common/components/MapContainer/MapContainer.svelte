<script lang="typescript">

    import { MapItemData } from '../MapActionElement/utils';

    import { yandexMapsReady, googleMapsReady } from '../../../stores/maps';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    import { onMount } from 'svelte';

    onMount(async () => {
        // const config = {
        //     name,
        //     addresses,
        // }
        const placemarkBuilder: PlacemarkBuilder = new YandexMapsPlacemarkBuilder();
        const mapBuilder: MapsBuilder = new YandexMapsBuilder(container, placemarkBuilder);

        mapBuilder.findMultipleAddressesAndCreateMap(data);
    })

    let container;

    export let data: MapItemData[] = [];

    export let addresses: string[] = [];
    export let name: string = "";
</script>

<div class="containerWrapper" style="width: 100%; height: 550px; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px;">
    {#if yandexMapsReady && googleMapsReady}
        <div style="width:100%; height: 100%;" bind:this={container}>
        </div>
    {/if}
</div>


