<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {Shop} from "../types";
    import {removeShopFromStore} from "../../../stores/shops";
    import {EntityType, ActionType} from "../../../stores/utils";

    import {style} from "./style";

    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte"
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte"

    let isPopupOpened = false;

    const onAcceptHandler = async () => {
        await removeShopFromStore(data.id);
    }

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: Shop = {};
</script>

<div>
    <ButtonIconRemove onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer entityType={EntityType.Shop}
                        actionType={ActionType.Remove} entityId={data.id}
                        onAcceptHandler={onAcceptHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={true}
                        title={`${$_("common.titles.remove")} ${data.name}` || $_("common.titles.remove")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        popupClass={style.Popup}>
            <div class="{style.SideMainPadding} {style.FlexHorCenter} {style.Message} {style.Font732Black}">
                {`${$_("shops.messages.remove_part1")}${data.name}${$_("shops.messages.remove_part2")}`}
            </div>
        </PopupContainer>
    {/if}
</div>