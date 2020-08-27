<script lang="typescript">

    import { MapItemData } from '../../ActionElements/MapActionElement/utils';

    import {style} from "./style";

    import { yandexMapsReady, googleMapsReady } from '../../../../stores/maps';

    import {onMount} from 'svelte';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    const processMapData = (data, container) => {
        if (data.length > 1) {
            mapBuilder && mapBuilder.findMultipleAddressesAndCreateMap(data, container);
        } else {
            mapBuilder && mapBuilder.findSingleAddressAndCreateMap(data[0], container);
        }
    }

    let container;

    onMount(() => {
        placemarkBuilder = new YandexMapsPlacemarkBuilder();
        mapBuilder = new YandexMapsBuilder(placemarkBuilder, isEditable, updateAddress);
    })

    let placemarkBuilder: PlacemarkBuilder = null;
    let mapBuilder: MapsBuilder = null;

    $: processMapData(data, container);

    export let data: MapItemData[] = [];
    export let wrapperClass: string = "";
    export let isEditable: boolean = false;
    export let updateAddress = () => {};
</script>

<div class="{style.ContainerWrapper} {wrapperClass}">
    {#if yandexMapsReady && googleMapsReady}
        <div class="{style.Container}" bind:this={container}>
        </div>
    {/if}
</div>


