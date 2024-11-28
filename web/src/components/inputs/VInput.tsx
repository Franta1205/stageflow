import PropTypes from "prop-types";

interface VinputProps {
    label?: string;
    type: 'email' | 'password';
}

export const VInput = (props: VinputProps) => {
    return (
        <div>
            {props.label && (<label htmlFor="email" className="block text-sm/6 font-medium text-gray-900">
                {props.label}
            </label>)}
            <div className="mt-2">
                <input
                    id={props.type}
                    name={props.type}
                    type={props.type}
                    required
                    autoComplete={props.type}
                    className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm/6"
                />
            </div>
        </div>
    )
}

VInput.propTypes = {
    label: PropTypes.string,
    type: PropTypes.string
}
