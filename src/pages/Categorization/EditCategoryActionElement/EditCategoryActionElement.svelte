<script lang="ts">
    import {onMount} from "svelte";
    import {Category} from "../types";
    import Button from "../../../common/components/Buttons/Button/Button.svelte";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {Popup} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";
    import CategoryEditForm from "../CategoryEditForm/CategoryEditForm.svelte";
    import {
        addCategoryToStore,
        updateCategoryInStore
    } from "../../../stores/categories";

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
            updateCategoryInStore(data)
        } else {
            addCategoryToStore(data)
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
                        withCancelButton="Cancel"
                        entityType={EntityType.Category}
                        actionType={ActionType.Remove} entityId={data.id}
                        title={data.id ? `Edit ${data.name}` || "Edit" : "New category"}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle="Save">
            <CategoryEditForm outerPopupUuid={uuid} data={data}/>
        </PopupContainer>
    {/if}

</div>