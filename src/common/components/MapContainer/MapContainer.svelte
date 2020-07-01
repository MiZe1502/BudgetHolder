<script lang="typescript">

    import { MapItemData } from '../MapActionElement/utils';

    import { ContainerWrapper, Container } from "./style";

    import { yandexMapsReady, googleMapsReady } from '../../../stores/maps';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    const processMapData = (data, container) => {
        if (mapBuilder.map) {
            return;
        }
        if (data.length > 1) {
            mapBuilder.findMultipleAddressesAndCreateMap(data, container);
        } else {
            mapBuilder.findSingleAddressAndCreateMap(data[0], container);
        }
    }

    let container;

    const placemarkBuilder: PlacemarkBuilder = new YandexMapsPlacemarkBuilder();
    const mapBuilder: MapsBuilder = new YandexMapsBuilder(placemarkBuilder);

    $: processMapData(data, container);

    export let data: MapItemData[] = [];
    export let wrapperClass: string = "";
</script>

<div class="{ContainerWrapper} {wrapperClass}">
    {#if yandexMapsReady && googleMapsReady}
        <div class="{Container}" bind:this={container}>
        </div>
    {/if}
</div>


