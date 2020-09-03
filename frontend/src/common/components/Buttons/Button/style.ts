import { css } from 'emotion';

export { Font312Black } from "../../../styles/fonts";
export { ButtonReset, ButtonDefault, ButtonSecondary, ButtonPrimary } from "../../../styles/buttons";

export const Button = css`
	width: auto;
	height: auto;
	padding: 8px;
	border-radius: 4px;
	box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.3);

	&:hover,
	&:active {
        box-shadow: 0 3px 1px -1px rgba(0, 84, 72, 0.15);
	}
	
	
	&[disabled]:hover {
	    box-shadow: none;
    }

	&:active {
        box-shadow: none;
	}
`;