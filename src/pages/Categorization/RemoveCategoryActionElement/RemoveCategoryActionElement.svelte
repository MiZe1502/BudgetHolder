<script lang="ts">
    import ButtonIconRemove
        from "../../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {Category} from "../types";
    import {removeCategoryFromStore} from "../../../stores/categories";
    import {EntityType, ActionType} from "../../../stores/utils";
    import {
        Popup,
        Font732Black,
        Message,
        SideMainPadding,
        FlexHorCenter
    } from "./style";

    let isPopupOpened = false;

    const onAcceptHandler = () => {
        removeCategoryFromStore(data.id);
    }

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
    }

    export let data: Category = {};

</script>

<!--TODO: implement action wrapper to prevent copy-paste-->
<div>
    <ButtonIconRemove width={16} height={16} onClickHandler={onClickHandler}/>
    {#if isPopupOpened}
        <PopupContainer entityType={EntityType.Category}
                        actionType={ActionType.Remove} entityId={data.id}
                        onAcceptHandler={onAcceptHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={true}
                        title={`Remove ${data.name}` || "Remove"}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler} popupClass={Popup}>
            <div class="{SideMainPadding} {FlexHorCenter} {Message} {Font732Black}">
                Do You really want to remove {data.name}? All children
                categories will be erased too.
            </div>
        </PopupContainer>
    {/if}
</div>