<script lang="typescript">
    import CommonActionElement from "../CommonActionElement/CommonActionElement.svelte";
    import MapActionElement from "../MapActionElement/MapActionElement.svelte";

    import { ShopActionsData} from "./utils";

    export let data: ShopActionsData = {};
</script>

<CommonActionElement>
    <MapActionElement data={data.mapData} />
</CommonActionElement>

