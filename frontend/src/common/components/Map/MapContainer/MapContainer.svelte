<script lang="typescript">
    import { get } from 'svelte/store'
    import { MapItemData } from '../../ActionElements/MapActionElement/utils';

    import {style} from "./style";

    import { yandexMapsReady, googleMapsReady } from '../../../../stores/maps';

    import {onMount} from 'svelte';

    import { YandexMapsBuilder } from "./YandexMapsBuilder";
    import { MapsBuilder } from "./MapsBuilder";
    import { PlacemarkBuilder } from "./PlacemarkBuilder";
    import { YandexMapsPlacemarkBuilder } from "./YandexMapsPlacemarkBuilder";

    const drawMap = (data, container) => {
        if (data.length > 1) {
            mapBuilder && mapBuilder.findMultipleAddressesAndCreateMap(data, container);
        } else {
            mapBuilder && mapBuilder.findSingleAddressAndCreateMap(data[0], container);
        }
    }

    const processMapData = (data, container) => {
        if (!get(yandexMapsReady)) {
            setTimeout(() => {
                drawMap(data, container)
            }, 500)
        } else {
            drawMap(data, container)
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


