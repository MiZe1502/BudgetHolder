<script lang="typescript">

    import { yandexMapsReady, googleMapsReady } from '../../../stores/maps';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    import { onMount } from 'svelte';

    onMount(async () => {
        const config = {
            name,
            address,
        }
        const placemarkBuilder: PlacemarkBuilder = new YandexMapsPlacemarkBuilder(config);
        const mapBuilder: MapsBuilder = new YandexMapsBuilder(container, placemarkBuilder);

        mapBuilder.findSingleAddressAndCreateMap(config.address);
    })

    let container;

    export let address: string = "";
    export let name: string = "";
</script>

<div class="containerWrapper" style="width: 100%; height: 550px; border-bottom-right-radius: 4px; border-bottom-left-radius: 4px;">
    {#if yandexMapsReady && googleMapsReady}
        <div style="width:100%; height: 100%;" bind:this={container}>
        </div>
    {/if}
</div>


