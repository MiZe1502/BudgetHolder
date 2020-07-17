<script lang="typescript">
    import {onMount} from "svelte";
    import {Shop} from "../../../pages/types";
    import {updateShopInStore, addShopToStore} from "../../../stores/shops";

    import {Popup} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";

    import ButtonIconEdit from "../Buttons/ButtonIconEdit/ButtonIconEdit.svelte"
    import PopupContainer from "../PopupContainer/PopupContainer.svelte"
    import ShopEditForm from "../ShopEditForm/ShopEditForm.svelte";
    import Button from "../Buttons/Button/Button.svelte";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = {...initialData}
    }

    const onSaveHandler = () => {
        if (data.id) {
            updateShopInStore(data);
        } else {
            addShopToStore(data);
        }
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: Shop = {};

    export let data: Shop = {}
    export let withButton: boolean = false;
    export let buttonTitle: string = "";
</script>

<div>
    {#if withButton}
        <Button title={buttonTitle} onClickHandler={onClickHandler}/>
    {:else}
        <ButtonIconEdit onClickHandler={onClickHandler}/>
    {/if}
    {#if isPopupOpened}
        <PopupContainer let:outerPopupUuid={uuid}
                        popupClass={Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton="Cancel" entityType={EntityType.Shop}
                        actionType={ActionType.Remove} entityId={data.id}
                        title={data.id ? `Edit ${data.name}` || "Edit" : "New shop"}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle="Save">
            <ShopEditForm outerPopupUuid={uuid} shop={data}/>
        </PopupContainer>
    {/if}
</div>