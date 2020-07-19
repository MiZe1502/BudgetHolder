<script lang="ts">
    import {onMount} from "svelte";
    import {Category} from "../types";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {Popup} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";

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
            console.log("update")
        } else {
            console.log("add")
        }
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: Category = {};

    export let data: Category = {};
    export let withButton: boolean = false;
    export let buttonTitle: string = "";
</script>


<div>
    {#if withButton}
        <Button title={buttonTitle} onClickHandler={onClickHandler}/>
    {:else}
        <ButtonIconEdit width={16} height={16} onClickHandler={onClickHandler}/>
    {/if}
    {#if isPopupOpened}
        <PopupContainer let:outerPopupUuid={uuid}
                        popupClass={Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton="Cancel" entityType={EntityType.Category}
                        actionType={ActionType.Remove} entityId={data.id}
                        title={data.id ? `Edit ${data.name}` || "Edit" : "New category"}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle="Save">
            Category form
            <!--            <ShopEditForm outerPopupUuid={uuid} shop={data}/>-->
        </PopupContainer>
    {/if}

</div>