<script lang="typescript">
    import {MapContainerWrapper, Popup} from "./style";

    import ButtonIconMap from "../../Buttons/ButtonIconMap/ButtonIconMap.svelte"
    import PopupContainer from "../../PopupContainer/PopupContainer.svelte"
    import MapContainer from "../../Map/MapContainer/MapContainer.svelte";

    import {onMount} from 'svelte';

    import {MapActionElementData, MapItemData} from "./utils";

    let isPopupOpened = false;

    let isDomReady = false;

    const onClickHandler = (event: MouseEvent) => {
        event.preventDefault();

        isPopupOpened = true;
        setTimeout(() => {
            isDomReady = true;
        }, 100)

    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: MapItemData[] = [];
    export let isEditable: boolean = false;
    export let updateAddress = () => {};
    export let onCancel = () => {};
</script>

<div>
    <ButtonIconMap onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer entityType={data && data[0].entityType}
                        actionType={data && data[0].actionType}
                        entityId={data && data[0].id}
                        title={data[0].name || "Map"}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        withAcceptButton={isEditable}
                        withCancelButton={isEditable}
                        onCancelHandler={onCancel}
                        popupClass={Popup}>
            {#if isDomReady}
                <MapContainer {isEditable} {data}
                              updateAddress={updateAddress}
                              wrapperClass={MapContainerWrapper}/>
            {/if}
        </PopupContainer>
    {/if}
</div>