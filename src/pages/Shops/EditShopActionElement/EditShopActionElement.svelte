<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {Shop} from "../types";
    import {updateShopInStore, addShopToStore} from "../../../stores/shops";

    import {style} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";

    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte"
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte"
    import ShopEditForm from "../ShopEditForm/ShopEditForm.svelte";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";

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
                        popupClass={style.Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={$_("common.components.buttons.cancel")}
                        entityType={EntityType.Shop}
                        actionType={ActionType.Edit} entityId={data.id}
                        title={data.id ? `${$_("common.titles.edit")} ${data.name}` || $_("common.titles.edit") : $_("shops.titles.new")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle={$_("common.components.buttons.save")}>
            <ShopEditForm outerPopupUuid={uuid} shop={data}/>
        </PopupContainer>
    {/if}
</div>